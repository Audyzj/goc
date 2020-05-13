/*
 Copyright 2020 Qiniu Cloud (七牛云)

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

package app

import (
	"github.com/qiniu/goc/pkg/cover"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start a server to host all services",
	Run: func(cmd *cobra.Command, args []string) {
		cover.StartServer(port)
	},
}

var port string

func init() {
	serverCmd.Flags().StringVarP(&port, "port", "", ":7777", "listen port to start a coverage host center")
	rootCmd.AddCommand(serverCmd)
}
