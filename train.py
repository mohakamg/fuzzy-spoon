# Import Libraries
import tensorflow as tf
from tensorflow import keras as tfk
import numpy as np
import matplotlib.pyplot as plt

# Constants
NUM_POINTS = 1000

# Create our Dataset
x = np.linspace(0, 3.142 * 2, NUM_POINTS)
y = np.sin(x) + np.random.normal(0, 0.2, size=NUM_POINTS)

# Plot
#plt.scatter(x, y)
#plt.show()

# Build the model
inp = tfk.layers.Input(shape=(1))
nonlinear = tfk.layers.Dense(10, activation="tanh")(inp)
op = tfk.layers.Dense(1, activation="linear", use_bias=True)(nonlinear)
model = tfk.models.Model(inp, op)
model.compile(optimizer="Adam", loss="mse", metrics=["mae"])
print(model.summary())


# Train the model
model.fit(
    x, y, epochs=200, verbose=False
)

# Generate Predictions
preds = model.predict(x)

# Save model
print(model.input.name)
print(model.output.name)
model.save("sineestimator/3")

# Plot Predictions
plt.scatter(x, y, c='r')
plt.scatter(x, preds, c='b')
plt.show()


