#!/bin/bash

# chmod +x gorm.sh
# ./gorm.sh proto/chatpb/chat.pb.go

g () {
  sed "s/json:\"$1,omitempty\"/json:\"$1,omitempty\" gorm:\"type:$2\"/"
}

cat $1 \
| g "id" "primary_key" \
| g "name" "varchar(100)" \
> $1.tmp && mv $1{.tmp,}