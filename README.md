# terraform-list-backends
List Status of Terraform States in Backend.

* This is the equivalent of the List Stacks command of Cloud Formation

* Currently only works with AWS S3 as the backend

## build
```
go build -o listbackends.exe
```
## Usage 
```
listbackends <S3 bucket>
```
e.g.
```
listbackends mystatebucket
```

## Output 
```
Terraform State                                                Status
---------------                                                ------
dev/example_app/terraform.state                                ACTIVE WITH RESOURCES
dev/app/terraform.state                                        ACTIVE WITH RESOURCES
dev/aws_scheduler/terraform.state                              ACTIVE WITH RESOURCES
dev/db/terraform.state                                         ACTIVE WITH RESOURCES
dev/domainapi/terraform.state                                  NO RESOURCES
dev/vpc/terraform.state                                        NO RESOURCES
```

## To Do

* Extend to other backends besides AWS S3
* Merge code in Terraform with a new command list-backends
