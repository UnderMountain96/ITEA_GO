Feature: Update article

  Scenario: Success
    When I run "update_article" command with params "id=6912354f-43b4-4106-8744-d84471adf59b,body=test"
    Then I see in "articles" table with params "id=6912354f-43b4-4106-8744-d84471adf59b":
      | body         |
      | test         |