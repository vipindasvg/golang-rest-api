# file: posts.feature
Feature: get posts
  I need to be able to request list of posts

  Scenario: does not allow PUT method
    When I send "PUT" request to "/posts"
    Then the response code should be 405

  Scenario: does not allow DELETE method
    When I send "DELETE" request to "/posts"
    Then the response code should be 405  

  Scenario: should create post record
    And I have following request body:
      """
      {"title": "Vikhil","text": "Das"}
      """
    When I send "POST" request to "/posts"
    Then the response code should be 200

  Scenario: should get lists of posts
    When I send "GET" request to "/posts"
    Then the response code should be 200