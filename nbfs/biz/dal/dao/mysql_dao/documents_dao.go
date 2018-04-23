/*
 *  Copyright (c) 2018, https://github.com/nebulaim
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

package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/nbfs/biz/dal/dataobject"
)

type DocumentsDAO struct {
	db *sqlx.DB
}

func NewDocumentsDAO(db *sqlx.DB) *DocumentsDAO {
	return &DocumentsDAO{db}
}

// insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id)
// TODO(@benqi): sqlmap
func (dao *DocumentsDAO) Insert(do *dataobject.DocumentsDO) int64 {
	var query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, version from documents where dc_id = 2 and document_id = :document_id and access_hash = :access_hash and version = :version
// TODO(@benqi): sqlmap
func (dao *DocumentsDAO) SelectByFileLocation(document_id int64, access_hash int64, version int32) *dataobject.DocumentsDO {
	var query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, version from documents where dc_id = 2 and document_id = ? and access_hash = ? and version = ?"
	rows, err := dao.db.Queryx(query, document_id, access_hash, version)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByFileLocation(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.DocumentsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByFileLocation(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
