# Coffee Machine

### Assumptions:

This project is more about the functional requirements specified in the problem statement.
Hence no api's are exposed and there is no single dependency involved.
1. The test format has to be consistent with the format provided [here]( ​https://www.npoint.io/docs/e8cd5a9bbd1331de326a).
2. Name of each beverage and ingredient has to be unique.
3. To add and test a new configuration, you can do two things:
    
    a. Go to `coffee_machine/test/behaviour/test_cases` and change the `test_case1.json`
    
    b. Else add a new file say `test_case2.json` at `coffee_machine/test/behaviour/test_cases` and change the file name in
    `coffee_machine/test/behaviour/base.go#56`
    
### Requirements to Run:

This project has no dependency. Only go setup is required. This is the version on my machine `go version go1.11.5 darwin/amd64`
A test run video name `test_run_video.mov` is present which actually shows the simulation of functionality.


### Code Structure
Top Level packages are pkg and test.
1. `pkg` is public package support recommended by go. Each subsequent package has file with same name to define the interface contract if required. 
    
    a. It has one `constant` package (required to define necessary constant).

    b. There is one `dataservice` package which handles the data CRUD operations. To eliminate external dependency, inmemory storage is preffered
and it has been tried to make api imitate db behaviour.

    c. `model` package contains bussiness dao objects. `DrinkRecipe` and `Ingredient` are used for storage
    while  `MachineOutletPool` is used to maintain machine state.
    
    d. `use_case` package is used to define the business logic.
    
2. `test` package contains two units i.e. behaviour and functional

    a. `behaviour` contains one test which tries to imitate the real world scenario. 10 random users will select one beverage at regular interval
    and ingredients will be refilled after certain amount of time. separate go routine are run to imitate the concurrency and parallelism use cases.
    
    c. `functional` contains basic functional testing of use_case units.
    

 Issue Reporting
 ----------------
 If you found any issues or wanted to discuss about any approach please shoot an email to bansalaman2905[at]gmail[dot]com.
