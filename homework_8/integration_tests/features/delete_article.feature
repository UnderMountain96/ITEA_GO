Feature: Delete article

  Scenario: Success
    When I run "delete_article" command with params "id=6912354f-43b4-4106-8744-d84471adf59b"
    Then I don't see record with ID "6912354f-43b4-4106-8744-d84471adf59b" in "articles" table:
