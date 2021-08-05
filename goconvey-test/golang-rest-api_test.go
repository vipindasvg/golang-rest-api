package main_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGolangRestApi(t *testing.T) {

	Convey("Subject: Testing Api Requests", t, func() {

		Convey("Given a requests for get posts", func() {

			Convey("When requests run successfully", func() {
				
				Convey("User gets the posts", nil)

			})

			Convey("When request not found", func() {

				Convey("User will not get the posts", nil)

			})

		})

	})

}