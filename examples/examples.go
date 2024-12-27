// Example Usages
// task-tracer add "task1" -> add a task named task1. Gives an ID and assigns status as "todo".
// task-tracker update 1 "new-task-name" -> If a task with ID 1 exists, updates the desciption with "new-task-name". If not, raises an error
// task-tracker delete 1 -> If a task with ID 1 exists, deletes it. If not, nothing happens and program ends with 0.
// task-tracker mark-in-progress 1 -> If a task with ID 1 exists, changes its status to "in-progress". If its status is already in-progress, prints a message.
// task-tracker mark-done 1 -> If a task with ID 1 exists, changes its status to "done". If its status is already done, prints a message.
// task-tracker list (todo,in-progresss,done)-> Lists all of the tasks. If any of the status values passed as a argument, filters and then lists the tasks.

package examples