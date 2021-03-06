package storage

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	grpc "github.com/galamshar/microservices-wallet/user/grpc/client"
	"github.com/galamshar/microservices-wallet/user/models"
)

//CheckExistingUser Check existing User
func (u *UserStorageService) CheckExistingUser(ID string) (bool, bool, error) {

	var userDB *models.User = new(models.User)

	var (
		exits    bool
		isActive bool
	)

	user, _ := u.GetUserCache(ID)

	if user != nil {
		exits = true
		isActive = true

		return exits, isActive, nil
	}

	if err := u.db.Where("user_id = ?", ID).First(userDB).Error; err != nil {
		exits = false
		isActive = false

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exits, isActive, errors.New("User Not exists")
		}
		return exits, isActive, err
	}

	exits = true

	if !userDB.IsActive {
		isActive = false
		return exits, isActive, nil
	}

	isActive = true

	return exits, isActive, nil
}

//GetIDName Get the ID from the username
func (u *UserStorageService) GetIDName(username string, email string) (string, error) {
	var userDB *models.User = new(models.User)

	if err := u.db.Where(&models.User{UserName: username}).First(&userDB).Error; err != nil {
		return "", err
	}

	if len(userDB.UserID.String()) > 0 || userDB.UserID.String() != "" {
		return userDB.UserID.String(), nil
	}

	var profileDB *models.Profile = new(models.Profile)

	if err := u.db.Where(&models.Profile{Email: email}).First(&profileDB).Error; err != nil {
		return "", err
	}

	return profileDB.UserID.String(), nil
}

//CheckExistingRelation Check if exits any relations before create
func (u *UserStorageService) CheckExistingRelation(fromUser string, toUsername string, active bool) (bool, error) {
	//Check values
	if len(fromUser) < 0 || len(toUsername) < 0 {
		return false, errors.New("Must send boths variables")
	}

	var relationDB *models.Relation = new(models.Relation)

	err := u.db.Where(&models.Relation{FromName: fromUser, ToName: toUsername}).
		Or(&models.Relation{FromName: toUsername, ToName: fromUser, Mutual: true}).First(&relationDB).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	//If pass the variable and the relation is not active, reactive the relation
	if active == true && relationDB.IsActive == false {
		//Update the DB
		err = u.db.Model(&models.Relation{}).Where(&models.Relation{RelationID: relationDB.RelationID, IsActive: false}).Update("is_active", true).Error
		if err != nil {
			return false, err
		}

		var wg sync.WaitGroup

		//Update the Cache
		wg.Add(2)
		go func() {
			u.UpdateRelations(relationDB.FromUser.String())
			wg.Done()
		}()
		go func() {
			u.UpdateRelations(relationDB.ToUser.String())
			wg.Done()
		}()

		wg.Wait()

		//Create the movement
		succes, err := grpc.CreateMovement("Relations", "Update relation to mutual", "User Service")

		if err != nil {
			log.Println("Error in Create a movement: " + err.Error())
		}

		if succes == false {
			log.Println("Error in Create a movement")
		}

		return true, errors.New("The relation was reactived")
	}

	return true, nil
}

//CheckMutualRelation Check if exits a relation and if is not mutual, If is not mutual update it
func (u *UserStorageService) CheckMutualRelation(fromUser string, fromID string, toUsername string) (bool, error) {
	//Check values
	if len(fromUser) < 0 || len(toUsername) < 0 {
		return false, errors.New("Must send boths variables")
	}

	var relationDB *models.Relation = new(models.Relation)

	//If the relations already exits with other user updated to mutual
	err := u.db.Where(&models.Relation{FromName: toUsername, ToName: fromUser, Mutual: false}).First(&relationDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return true, err
	}

	if relationDB.Mutual == false {
		err := u.db.Model(&models.Relation{}).
			Where(&models.Relation{RelationID: relationDB.RelationID}).
			Updates(map[string]interface{}{"mutual": true, "updated_at": time.Now()}).Error
		if err != nil {
			return false, err
		}

		if err := u.UpdateRelations(fromID); err != nil {
			return true, err
		}

		succes, err := grpc.CreateMovement("Relations", "Update relation to mutual", "User Service")

		if err != nil {
			log.Println("Error in Create a movement: " + err.Error())
		}

		if succes == false {
			log.Println("Error in Create a movement")
		}

		toID, err := u.GetIDName(toUsername, "")

		if err != nil {
			log.Println(toUsername)
			return false, err
		}

		if toID != "" {
			u.UpdateRelations(toID)
		}

		return true, nil
	}

	if relationDB.FromName != "" && relationDB.FromName == fromUser {
		return true, nil
	}
	return false, nil
}
