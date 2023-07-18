#!/bin/sh
workDir="/home/go/go/src/apis/"
dockerWorkDir="/main"
binName="apis"
projectName="go-apis"
version=${1}
dbFile=${2}

dockerContainerPort="50001"
dockerContainerName="go-apis-"${dockerContainerPort}


timeStr=`date +%Y%m%d`
imageVersion=${timeStr}_${version}

binDir=${workDir}"bin/"
scriptDir=${workDir}"sh/"
dockerfileTemp=${scriptDir}Dockerfile

dockerfileDir="/server/"${projectName}"/"
dockerFile=${dockerfileDir}Dockerfile
dockerdataDir="/data/"${projectName}"_docker/"
alias sdocker="sudo docker"
oldImage=`sdocker ps | grep ${dockerContainerName} | awk '{print $2}'`

# ask
echo "get imageVersion "${imageVersion}

echo "get dbFile "${dbFile}

echo "i can't update conf, are you update it??? " 

read -p "update? (y/n): " input

if [ "$input" = "y" ]; then
    echo "updating ..."
elif [ "$input" = "n" ]; then
    echo "exiting ..."
    exit 0
else
    echo "input err"
    exit 1
fi

if [[ -z "${dbFile}" ]]; then
    echo "no db source"
else
    # bak db
    mkdir -p ${dockerdataDir}dbbak_${imageVersion}
    cp ${dockerdataDir}db/* ${dockerdataDir}/dbbak_${imageVersion}
    # source sql to db
    sqlite3 ${dockerdataDir}db/apis.db < ${workDir}sql/${dbFile}
fi


# go build
cd ${workDir}
mkdir -p ${binDir}
go build -o ${binDir}

# copy dockerfile
rm -f ${dockerfileTemp}
echo "FROM ubuntu:with_bag" >> ${dockerfileTemp}
echo "WORKDIR "${dockerWorkDir} >> ${dockerfileTemp}
echo "COPY . ." >> ${dockerfileTemp}
echo "ENV TZ=Asia/Shanghai" >> ${dockerfileTemp}
echo "CMD [\""${dockerWorkDir}"/"${binName}"\"]" >> ${dockerfileTemp}

rm ${dockerFile}
cp -pr ${dockerfileTemp} ${dockerFile}

# docker build
cd ${dockerfileDir}
rm -f ${dockerfileDir}${binName}
cp ${binDir}${binName} ${dockerfileDir}
sdocker build . -t ${projectName}:${imageVersion}
 

# docker deploy
sdocker rm -f ${dockerContainerName}

sdocker run -d --name ${dockerContainerName} -p ${dockerContainerPort}:${dockerContainerPort} --restart=always -v ${dockerdataDir}db:${dockerWorkDir}/db -v ${dockerdataDir}config:${dockerWorkDir}/config -v ${dockerdataDir}log:${dockerWorkDir}/log  ${projectName}:${imageVersion}
