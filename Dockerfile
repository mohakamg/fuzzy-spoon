# Inherit from a Base Image
FROM tensorflow/serving

# Copy the Model
COPY ./sineestimator /models/sineestimator

ENV MODEL_NAME sineestimator
