/*
 * Copyright (C) 2024 Key9 Identity, Inc <k9.io>
 * Copyright (C) 2024 Champ Clark III <cclark@k9.io>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License Version 2 as
 * published by the Free Software Foundation.  You may not use, modify or
 * distribute this program under any other version of the GNU General
 * Public License.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
 */

package main

import (
        "bytes"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"

        "github.com/gin-gonic/gin"
)

func Process_Client_Log(c *gin.Context) {

        var req *http.Request
        var err error
        var jsondata []uint8

	jsondata, err = c.GetRawData()

	if err != nil || jsondata == nil {
		log.Printf("Can't process JSON data from client.\n")
                c.JSON(http.StatusOK, gin.H{"error": "Error processing POST JSON data"})
                c.Abort()
                return
	}

        client := http.Client{}

        api_key_temp := fmt.Sprintf("%s:%s", c.GetString("company_uuid"), c.GetString("api_key"))
        url_tmp := fmt.Sprintf("%s%s", Config.Core.Address, c.Request.URL.Path)

        req, err = http.NewRequest("POST", url_tmp, bytes.NewBuffer(jsondata))

	if err != nil {

		log.Printf("Can't build new proxy request [http.NewRequest]\n")
		c.JSON(http.StatusOK, gin.H{"error": "Can't build new proxy request"})
		c.Abort()
		return

	}

	req.Header["API_KEY"] = []string{api_key_temp}

	res, err := client.Do(req)

	if err != nil {

		log.Printf("Can't build new proxy request [client.Do]\n")
		c.JSON(http.StatusOK, gin.H{"error": "Can't build new proxy request"})
		c.Abort()
		return

	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil { 

		log.Printf("Can't read body [ioutil.ReadAll]\n")
		c.JSON(http.StatusOK, gin.H{"error": "Read data from proxy request"})
		c.Abort()
		return

	}

	/* Send data to the client */

	c.Data(http.StatusOK, "application/json", body)

}
