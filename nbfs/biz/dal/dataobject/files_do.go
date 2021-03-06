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

package dataobject

type FilesDO struct {
	Id            int64  `db:"id"`
	FileId        int64  `db:"file_id"`
	AccessHash    int64  `db:"access_hash"`
	CreatorId     int64  `db:"creator_id"`
	CreatorUserId int32  `db:"creator_user_id"`
	FilePartId    int64  `db:"file_part_id"`
	FileParts     int32  `db:"file_parts"`
	FileSize      int64  `db:"file_size"`
	FilePath      string `db:"file_path"`
	Ext           string `db:"ext"`
	IsBigFile     int8   `db:"is_big_file"`
	Md5Checksum   string `db:"md5_checksum"`
	UploadName    string `db:"upload_name"`
	CreatedAt     string `db:"created_at"`
}
