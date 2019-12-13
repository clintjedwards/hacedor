- Make the output of buildPlugin available in error message
- Flesh out RunPipeline now that it works
- We should clean up a build if the create fails
- Implement concurrent safe map for plugins and generating pipeline runs/or database these things..probably database
- Clean up a build if the create fails
- Tasks should be able to share children and that child should wait on both those tasks to complete successfully before moving on
- Tasks should support arguments
- Tasks should pass in some always availabile variables that would be useful to task creators, like which taskRunid is currently running
- Tasks should have some way to involve users manually. It would be cool for something to get to a state to wait for user interaction
- Tasks should be able to be executed on their lonesome. So I should be able to through the interface return a single task
- Check for cyclical graphs/ include check for if the root node even exits
- Check for any tasks that are in the map but are unable to be reached
- All these checks should fail the run and return a proper error so that that user can fix it
- Tasks should eventually properly show us logs
- Investigate using value passes instead of reference passes for ExecuteTask and friends
- Save plugins as they are added and reload them from the database on application startup
- CLI should use cool spinners because shutup
- ExecuteTask requests should have a parent status and run ID passed to them so that sdk users can use these vars to do interesting things. This means we need to implement the functionality of embedded arguments to the functions? Can we expose these another way?