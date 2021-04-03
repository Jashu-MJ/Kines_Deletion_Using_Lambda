# Kines_Deletion_Using_Lambda
By making a post call to AWS API gateway and by using lambda function, I have deleted the kinesis data stream



WINDOWS:

$env:GOOS = "linux"
go build -o main apiRequest.go
build-lambda-zip.exe -output kinesisC.zip main
