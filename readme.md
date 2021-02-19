```bash
# Parser (1 seul binaire)
$ ./parser input_google_BIG_DAY.txt
generate input_google_BIG_DAY.bin


# Solutionneurs (chacun sa sauce)
$ ./solutionneur1 input_google_BIG_DAY.bin
generate output_google_BIG_DAY_fromsoolutioonneur1.bin
$ python solutionneur2.py input_google_BIG_DAY.bin
generate output_google_BIG_DAY_fromsoolutioonneur2.bin
$ java solutionneur3.java input_google_BIG_DAY.bin
generate output_google_BIG_DAY_fromsoolutioonneur3.bin


# Optimizers (chacun sa sauce)
$ ./optimizer1 output_google_BIG_DAY_fromsoolutioonneur1.bin
generate output_google_BIG_DAY_fromsoolutioonneur1_optimized1.bin

# OutputGenerator
$ ./moulinette output_google_BIG_DAY_fromsoolutioonneur1_optimized1.bin
generate output_google_BIG_DAY_fromsoolutioonneur1_optimized1.txt


```