# Map reduce Implementation using wordcount in single machine in-memory version

This is using zero distributed systems complexity 

This is stage 1 of wordcount it doesn't use any goroutines or multicomputer capability or mutithreading , it just counts and stores the mappeddata and reduceddata 
in global variable :

The steps involved are :

1. It reads data from a file `file1.txt`
2. Counts the words
3. Maps the word in a Map called `MappedData`.
4. Then uses sorting algorithm to sort the `MappedData`.
5. Then use `reducer` function to reduce data into `ReducedData`.