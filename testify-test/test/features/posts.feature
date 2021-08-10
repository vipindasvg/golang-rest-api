# file: posts.feature
Feature: Test Golang rest apis
  I need to test all golang rest apis

  Scenario: does not allow PUT method
    When I send "PUT" request to "/posts"
    Then the response code should be 405

  Scenario: does not allow DELETE method
    When I send "DELETE" request to "/posts"
    Then the response code should be 405  

  Scenario: should create post record
    Given I have following request body:
      """
      {"title": "Vikhil","text": "Das"}
      """
    When I send "POST" request to "/posts"
    Then the response code should be 200

  Scenario: should get lists of posts
    When I send "GET" request to "/posts"
    Then the response code should be 200
    