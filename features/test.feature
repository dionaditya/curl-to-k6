Feature: create new user
  Scenario: Generate k6 script for create nwe user
    Given there is curl command:
    """
    curl -v -F filename=image.mp4 -F upload=@image.ppt http://localhost:8080/api/upload
    """
    And The script have stages options:
      """
        {
          stages: [
            {
                 duration: "5m", target: 60
            },
          ],
        }
      """
    And The file name is "generated_test.js"