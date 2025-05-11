package user_service

import (
	"context"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) UpdateUserTx(ctx context.Context, user *user.User, updateMask *fieldmaskpb.FieldMask) (*user.User, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if _, err := f.storage.UpdateUser(ctxTx, user, updateMask); err != nil {
			return errors.WrapFail(err, "update user")
		}

		if updateMask == nil || hasPath(updateMask.GetPaths(), "photo_upload_requests") {
			for _, photo := range user.GetPhotos() {
				if err := f.storage.UpdatePhoto(ctxTx, photo, user.GetID()); err != nil {
					return errors.WrapFail(err, "update photo")
				}
			}
		}

		if updateMask == nil || hasPath(updateMask.GetPaths(), "interest") {
			if err := f.storage.DeleteInterests(ctxTx, user.GetID()); err != nil {
				return errors.WrapFail(err, "delete old interests")
			}

			if err := f.storage.AddInterests(ctxTx, user.GetID(), user.GetInterest()); err != nil {
				return errors.WrapFail(err, "add new interests")
			}
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "UpdateUserTx failed")
	}
	return user, nil
}

func hasPath(paths []string, target string) bool {
	for _, path := range paths {
		if path == target {
			return true
		}
	}
	return false
}
