#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

PSQ_PACKAGE="postgresql"

PSQL_MINIMUM="9.4"

# will compare version using sort -V

check_version() { test "$(echo "$@" | tr " " "\n" | sort -V | tail -n 1)" == "$1"; }

add_repo() {
	echo "deb http://apt.postgresql.org/pub/repos/apt/ trusty-pgdg main" > /etc/apt/sources.list.d/pgdg.list
	wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
}


perform_install() {
	# On 14.04 systems, add the official repository first
	if [ $(lsb_release -rs) == "14.04" ]; then
		echo "Adding official PostgreSQL repository and installing ..."
		add_repo
	fi
	apt-get update
	apt-get --force-yes -y install $PSQ_PACKAGE
	if [ $? -ne 0 ]; then
		echo "Unable to install PostgreSQL"
		exit 1
	fi
}

PSQL_DPKG=`dpkg -s $PSQ_PACKAGE`
if [ $? -ne 0 ]; then
	perform_install
else
	# check if psql was removed
	STATUS=`dpkg -s  $PSQ_PACKAGE | grep Status | awk '{ print $2 }'`
	if [ $STATUS == "deinstall" ]; then
		perform_install
	else
		# The system does have postgresql installed, check the version
		PSQL_VERSION=`echo $PSQL_DPKG | grep Version | awk '{ print $2}'`
		if check_version $PSQL_VERSION $PSQL_MINIMUM; then
			echo "PostgreSQL is OK, continuing ..."
		else
			echo "PostgreSQL needs to be upgraded"
			perform_install

		fi
	fi
fi
# Database configuration
echo  "Configuring PostgreSQL user 'chef' ..."
sudo -u postgres psql postgres <<EOF
create user chef;
\password chef
create database cookbook with owner=chef;
EOF
echo -e "\nSetup completed. Make sure you update config.json to reflect the new database changes"

