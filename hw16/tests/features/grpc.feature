# file: features/web.feature


Feature: gRPC api

  Scenario: Add event in empty DB
    When I send execute AddEvent on "localhost:8889" Event data:
		"""
		{
			"Id": 1,
			"StartTime": 1000,
			"EndTime": 2000,
			"Title": "title1",
			"Description": "description1"
		}
		"""
    Then The response error code should be nil
    And The execute GetAllEvents return one same Event