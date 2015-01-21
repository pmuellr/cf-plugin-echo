# Licensed under the Apache License. See footer for details.

all: help

#-------------------------------------------------------------------------------
help:
	@echo "make targets:"
	@echo "  build:            build the plugin"
	@echo "  build-install:    run build, then install-plugin"
	@echo "  install-plugin:   re-install the plugin"
	@echo "  install-packages: install pre-req packages"
	@echo " "
	@echo "You may need to build 'install-packages' before 'build' is run."

#-------------------------------------------------------------------------------
build:
	go build cf-plugin-echo.go

#-------------------------------------------------------------------------------
build-install: build install-plugin

#-------------------------------------------------------------------------------
install-plugin:
	-cf uninstall-plugin echo
	cf install-plugin   cf-plugin-echo

#-------------------------------------------------------------------------------
install-packages:
	go get github.com/cloudfoundry/cli/plugin

#-------------------------------------------------------------------------------
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#-------------------------------------------------------------------------------
