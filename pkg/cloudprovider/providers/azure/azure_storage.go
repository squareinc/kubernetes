/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"fmt"

	"github.com/golang/glog"
)

const (
	defaultStorageAccountType      = string(storage.StandardLRS)
	fileShareAccountNamePrefix     = "f"
	sharedDiskAccountNamePrefix    = "ds"
	dedicatedDiskAccountNamePrefix = "dd"
)

// CreateFileShare creates a file share, using a matching storage account
func (az *Cloud) CreateFileShare(name, storageAccount, storageType, location string, requestGB int) (string, string, error) {
	account, key, err := az.ensureStorageAccount(accountName, accountType, location, fileShareAccountNamePrefix)
	if err != nil {
		return "", "", fmt.Errorf("could not get storage key for storage account %s: %v", accountName, err)
	}

	if err := az.createFileShare(account, key, shareName, requestGiB); err != nil {
		return "", "", fmt.Errorf("failed to create share %s in account %s: %v", shareName, account, err)
	}
	glog.V(4).Infof("created share %s in account %s", shareName, account)
	return account, key, nil
}

// DeleteFileShare deletes a file share using storage account name and key
func (az *Cloud) DeleteFileShare(accountName, key, name string) error {
	err := az.deleteFileShare(accountName, key, name)
	if err != nil {
		return err
	}
	glog.V(4).Infof("share %s deleted", name)
	return nil

}
