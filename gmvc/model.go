// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gmvc

import "github.com/kotlin2018/orm"

type (
	M     = Model      // M is alias for Model, just for short write purpose.
	Model = *orm.Model // Model is alias for *gdb.Model.
)
