// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HandlerFunc func(c *gin.Context) error

func WrapHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			var serverErr *Error
			if !errors.As(err, &serverErr) {
				serverErr = &Error{
					HTTPStatusCode: http.StatusInternalServerError,
					Err:            err,
					Message:        err.Error(),
				}
			}
			serverErr.Request = c.Request.Method + " " + c.Request.URL.String()

			if gin.Mode() == "debug" {
				logrus.Debugf("error: %+v", err)
			} else {
				// Remove detailed info when in the release mode
				serverErr.Op = ""
				serverErr.Err = nil
			}

			c.JSON(serverErr.HTTPStatusCode, serverErr)
			return
		}
	}
}
