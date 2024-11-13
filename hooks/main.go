package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

type Loger struct {
	logrus *logrus.Logger
}

func (l *Loger) LogMode(level logger.LogLevel) logger.Interface {
	l.logrus.SetLevel(logrus.PanicLevel)
	return l
}

func (l *Loger) Info(_ context.Context, a string, val ...interface{}) {
	l.logrus.Infof(a, val...)
}

func (l *Loger) Warn(_ context.Context, a string, val ...interface{}) {
	l.logrus.Warnf(a, val...)
}

func (l *Loger) Error(_ context.Context, a string, val ...interface{}) {
	l.logrus.Errorf(a, val...)
}

func (l *Loger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	entry := l.logrus.WithFields(logrus.Fields{
		"duration": time.Since(begin),
		"rows":     rows,
		"sql":      sql,
		"error":    err,
	})
	if err != nil {
		entry.Error("Trace error")
	} else {
		entry.Info("Trace success")
	}
}

type User struct {
	gorm.Model
	Name  string
	Email string
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("Before!")
	if user.Name == "" {
		return errors.New("Name can not be blank")
	}
	return nil
}

func main() {
	log := logrus.New()
	loger := &Loger{logrus: log}

	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: loger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to automigrate")
	}
	user := User{Name: "", Email: "sejo@email.com"}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to create user")
	}
}
