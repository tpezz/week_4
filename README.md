

Dataset: Kaggle Data: Anscombe's quartet (https://www.kaggle.com/datasets/carlmcbrideellis/data-anscombes-quartet?resource=download)

| x123 |  y1  |  y2  |  y3   | x4  |  y4  |
|------|------|------|------|-----|------|
| 10.0 | 8.04 | 9.14 | 7.46 | 8.0 | 6.58 |
|  8.0 | 6.95 | 8.14 | 6.77 | 8.0 | 5.76 |
| 13.0 | 7.58 | 8.74 |12.74 | 8.0 | 7.71 |
|  9.0 | 8.81 | 8.77 | 7.11 | 8.0 | 8.84 |
| 11.0 | 8.33 | 9.26 | 7.81 | 8.0 | 8.47 |


Python Regression Results:

The slopes and intercepts are very close to the expected values:
Expected Slope ≈ 0.5
Expected Intercept ≈ 3.0
The small variations (~0.0001 to 0.0025) are likely due to floating-point precision differences in computation.
Execution Time:

The reported execution times (~0.00004 to 0.00026 sec) are expected for a small dataset like this.

Dataset 1: Slope = 0.50009, Intercept = 3.00009, Time = 0.00026
Dataset 2: Slope = 0.50000, Intercept = 3.00091, Time = 0.00005
Dataset 3: Slope = 0.49973, Intercept = 3.00245, Time = 0.00004
Dataset 4: Slope = 0.49991, Intercept = 3.00173, Time = 0.00004