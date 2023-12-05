Feature: Update article

  Scenario: Success
    When I run "update_article" command with params "id=4ddc8d46-f08f-43da-b227-3afd79c69d16,body=test"
    Then I see record with ID "4ddc8d46-f08f-43da-b227-3afd79c69d16" in "articles" table: 
      | body         |
      | test         |
