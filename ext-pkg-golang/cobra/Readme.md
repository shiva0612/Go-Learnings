cmd hirarchy :

roots:
    items
        shirts
        shoes
    users

you will learn about:
1. how to make cmd hirachy using cobra
2. how to use persitent flags and local flags 

---NOTES---
go mod init test 
cobra-cli init test 
mv go.mod go.sum test/

cobra-cli add items
cobra-cli add shoes
cobra-cli add shirts


todo :

rootCmd
    users
    items   
        shirts
        shoes 

gm -h
    lists "items,user"

gm items -h
    lists "shoes,shirts"

gm items shirts -h 
    lists description of shirts 


---
in order to make the above hirarchy
just create the files using "cobra-cli add items, cobra-cli add users, cobra-cli add shirts, cobra-cli add shirts"
go to init() of shirts, shoes and change as mentioned below
rootCmd.AddCommand(shoesCmd) -> itemsCmd.AddCommand(shoesCmd)
lly, for shirts 
this creates the hirarchy
---

if you want few flags from parent to be used in its child cmds
simply add "persistentFlags" in the init() of the parent cmd

---

gm items shirts --instock=true --size=l puma nike
    flags are "instock, size"
    args are "puma, nike"

---