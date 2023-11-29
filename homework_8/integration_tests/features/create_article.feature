Feature: Create article

  Scenario: Success
    When I run "create_article" command
    Then I see in "articles" table:
      | title         |
      | This is title |
