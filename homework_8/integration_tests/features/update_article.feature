Feature: Update article

  Scenario: Update few fields
    When I run "update_article" command with params "id=a462db9b-b7ae-434c-87af-943d080d5c00,body=test,title=test"
    Then I see record with ID "a462db9b-b7ae-434c-87af-943d080d5c00" in "articles" table: 
      | title | body |
      | test  | test |
