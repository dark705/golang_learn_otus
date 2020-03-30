# file: features/web.feature


Feature: gRPC api

  Scenario: Execute AddEvent and then GetEvent in empty DB, return same event, after DelEvent
    When I send execute AddEvent on "localhost:8889" with data:
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
    And The execute GetEvent return same Event, with no error code
    And After Execute DelEvent, return no error code

  Scenario: Execute AddEvent two times on same event return error of busy day, and in Calendar only one event
    When I send execute AddEvent on "localhost:8889" two times with data:
		"""
		{
			"StartTime": 1000,
			"EndTime": 2000,
			"Title": "title1",
			"Description": "description1"
		}
		"""
    Then The response error desc should be "Date interval for new event already busy"