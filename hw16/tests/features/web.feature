# file: features/web.feature

# http://calendar_api:8081/
# http://localhost:8888/

Feature: Web api

	Scenario: Web api is available
		When I send "GET" request to "http://localhost:8888/"
		Then The response code should be 200
		And The response should match text "Hello world"
