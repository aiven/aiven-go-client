# Changelog

## [1.0.0] - 2018-09-26

Add support for all service types, service integrations, databases,
connection pools and VPCs. Also number of improvements to previously
supported APIs.

## [1.0.1] - 2019-05-29

Add support for specifying the API url via AIVEN_WEB_URL environment
variable

## [1.1.0] - 2019-06-06

Add support for cross region vpc peering.

## [1.3.0] - 2019-12-23

Add support for Elasticsearch ACLs.

## [1.4.0] - 2020-01-03

Add support for Kafka Schemas and Kafka Connectors

## [1.5.0] - 2020-02-04

Add support for Accounts:
 - accounts
 - account teams
 - account team members 
 - account team projects
 - account authentications 
 - account team invites  
 - projects add account id

Minor fixes for Kafka Schemas  

## [1.5.1] - 2020-02-07

Fix project creation and update which was depended on account_id property

## [1.5.2] - 2020-02-12

Add delete method for Account Team Invitations

## [1.5.3] - 2020-02-12

Add new SAML specific properties for Account Authentications

## [1.5.4] - 2020-05-07

Add support for Kafka Mirror Maker 2 Replication Flows

## [1.5.5] - 2020-05-12

Fix Kafka Mirror Maker 2 Replication Flows update functionality

## [1.5.6] - 2020-07-20

Add support for AWS Transit Gateway and expose Azure config parameters

## [1.5.7] - 2020-07-20
- Add service `powered` field to API requests and response
- Add editing possibility for Kafka Schema Registry subject
- Extend service components supported fields
- Add root CA support
- Use golang 1.14 Travis CI

## [1.5.8] - 2020-09-03
- Change service component ssl field to boolean
- Change accounts acceptance tests email
- Add Kafka schema subject configuration management
- Project related improvements: add new fields, helper functions and extend acceptance test
- Add Aiven error IsNotFound validation

## [1.5.9] - 2020-09-29
- Use golang 1.15
- Add PUT endpoint to modify service users
- Add/Get and Delete methods specific for Azure VPC peering connection
- Make Kafka Topic retention hours an optional field

## [1.5.10] - 2020-10-15
- Add support for Kafka Topic Config

