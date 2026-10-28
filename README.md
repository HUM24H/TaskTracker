# Task Tracker CLI - Built in Go
## Introduction
Task Tracker is a CLI app built using Go

**Note: This CLI app was built for the purpose of learning Go. This is not meant for any production or public use**

## Description
The Task Tracker CLI app is a simple Go application that allows a user to create a Task List.

Tasks within this list are stored within a JSON file and they can be provided with different statuses:
- pending
- in-progress
- done

The CLI app will also allow the user to return any number of items from the list based on their status value and it can also return the entire list.

## How to Use
In order to operate this CLI app, you will need to run the "TaskTracker" executable on a Linux machine.

Below are a list of the input commands that can be given to the executable:
- list --> Takes an additional argument which can be one of the following; done, todo, in-progress or all. This will return all the tasks where the status matches the argument passed in, with the exception of "all" which returns all tasks from the task.json file.
- add --> Takes an additional argument which is the Task Name and create a new entry in the task.json file.
- update --> Takes two additional arguments, the first additional argument is the Task ID which can be obtained by using the "list" function. The second argument is the new Task Name. This will then update the Task Name for the specified Task ID.
- update-status --> Takes two additional arguments, the first additional argument is the Task ID which can be obtained by using the "list" function. The second argument is the new Task Status, this can be one of the following; done, todo or in-progress" This will then update the Task Status for the specified Task ID.
- delete --> Takes an additional argument which is the Task ID, this can be obtained by using the "list function". This will then remove the specified Task ID from the task.json file.

Initial use will be to add a task, this can be done by running the following command:

`./TaskTracker add "Check Tyre Pressures"`

This will return a response to the CLI, below is the response when adding a task for the first time:

`JSON file does not exist.
JSON file created successfully.
Task added successfully.`

Below will be returned when adding any further tasks:

`JSON file exists.
Task added successfully.`

**Note: Any handled & unhandled errors from the application will also appear in the CLI as response**

From here you can now view the task list with the following command:

`./TaskTracker list all`

This will return the contents of the "tasks.json" file that is created and used to store the tasks:

`JSON file exists.
ID: 1, Name: Check Tyre Pressures, Status: pending`

Next we can update the existing task with a new status, for example we will update this to say it is now done by running the following command:

`./TaskTracker update-status 1 done`

This will then return the following:

`JSON file exists.
Task status updated successfully.`

From here we can then run the list command again and see the status of the task has now changed from "pending" to "done":

`JSON file exists.
ID: 1, Name: Check Tyre Pressures, Status: done`

If we wanted to change the name of the task, we can do that as well by running the following command:

`/TaskTracker update 1 "Check Type Pressures - Most important are front tyres"`

Then you will get the following response:

`JSON file exists.
Task updated successfully.`

Lastly we can remove a task by simply running the following command:

`./TaskTracker delete 1`

This will then return the following:

`JSON file exists.
Task deleted successfully.`
