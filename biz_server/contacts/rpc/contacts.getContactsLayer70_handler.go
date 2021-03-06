/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/logger"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz/core/user"
	"github.com/nebulaim/telegramd/biz/core/contact"
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

// @benqi: Android client
// contacts.getContactsLayer70#22c6aa08 hash:string = contacts.Contacts;
func (s *ContactsServiceImpl) ContactsGetContactsLayer70(ctx context.Context, request *mtproto.TLContactsGetContactsLayer70) (*mtproto.Contacts_Contacts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getContactsLayer70#22c6aa08 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		contacts *mtproto.Contacts_Contacts
	)
	contactLogic := contact.MakeContactLogic(md.UserId)

	contactList := contactLogic.GetContactList()
	// 避免查询数据库时IN()条件为empty
	if len(contactList) > 0 {
		// hash := ""
		md5Ctx := md5.New()
		for _, c := range contactList {
			md5Ctx.Write([]byte(fmt.Sprintf("%d", c.Id)))
		}
		cipherStr := md5Ctx.Sum(nil)
		md5hash := hex.EncodeToString(cipherStr)

		if md5hash == request.Hash {
			contacts = mtproto.NewTLContactsContactsNotModified().To_Contacts_Contacts()
		} else {
			idList := make([]int32, 0, len(contactList))
			cList := make([]*mtproto.Contact, 0, len(contactList))
			for _, c := range contactList {
				idList = append(idList, c.Id)
				c2 := &mtproto.Contact{
					Constructor: mtproto.TLConstructor_CRC32_contact,
					Data2: &mtproto.Contact_Data{
						UserId: c.Id,
						Mutual: mtproto.ToBool(c.Mutual == 1),
					},
				}
				cList = append(cList, c2)
			}

			users := user.GetUsersBySelfAndIDList(md.UserId, idList)
			contacts = &mtproto.Contacts_Contacts{
				Constructor: mtproto.TLConstructor_CRC32_contacts_contacts,
				Data2: &mtproto.Contacts_Contacts_Data{
					Contacts:   cList,
					SavedCount: int32(len(cList)),
					Users:      users,
				},
			}
		}
	} else {
		contacts = mtproto.NewTLContactsContacts().To_Contacts_Contacts()
	}

	glog.Infof("contacts.getContactsLayer70#22c6aa08 - reply: %s\n", logger.JsonDebugData(contacts))
	return contacts, nil
}
