# University info service
The service provides three different endpoints

## .../unisearcher/v1/uniinfo/
This endpoint retrieves university information from the foreign api's. 
The service then uses the retrieved information to combine the two api's.
The endpoint only accepts GET requests. Any other request will be received with
501 Not implemented status code. If no university was found it will re

## .../unisearcher/v1/neighbourunis/

<!-- #### Using the endpoint
Method: GET

Path: .../{:partial_or_complete_university_name}/

Example: .../norwegian%20university%20of%20science%20and%20technology/ -->






## .../unisearcher/v1/diag/