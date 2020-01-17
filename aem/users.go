// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import "context"

// UsersService handles communication with the Users related methods of Adobe
// AEM.
type UsersService service

func (s *UsersService) CreateUser(ctx context.Context) error {
	// curl -u admin:admin -FcreateUser= -FauthorizableId=testuser -Frep:password=abc123 http://localhost:4502/libs/granite/security/post/authorizables

	return _, nil
}

func (s *UsersService) ChangePassword(ctx context.Context) error {
	// curl -u admin:admin -F rep:password="test" http://localhost:4502/home/users/a/alister@geometrixx.com

	return _, nil
}
