# file: features/scheduler.feature

Feature: Scheduler feature

  Scenario: Scheduler not handle inappropriate events
    When I add Event in DB, with data:
		"""
		{
			"StartTime": "2000-01-02T15:00:00+03:00",
			"EndTime": "2000-01-02T16:00:00+03:00",
			"Title": "title1",
			"Description": "description1"
		}
		"""
    Then Then I wait "15" seconds to make sure that the scheduler processed event
    And Check DB for not scheduled events
    And Check event on Sender
    And Delete test event

  Scenario: Scheduler handle appropriate events
    When I add Event in DB, with data:
		"""
		{
			"StartTime": "2999-01-02T15:00:00+03:00",
			"EndTime": "2999-01-02T16:00:00+03:00",
			"Title": "title1",
			"Description": "description1"
		}
		"""
    Then Then I wait "15" seconds to make sure that the scheduler processed event
    And Check DB for scheduled events
    And Check no event on Sender
    And Delete test event
