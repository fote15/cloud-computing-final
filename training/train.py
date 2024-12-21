import numpy as np
from sklearn.linear_model import LinearRegression
import matplotlib.pyplot as plt

# Read the file content (similar to earlier example)
file_path = "data.txt"
with open(file_path, 'r') as file:
    file_content = file.read()

# Example: Extract Sales Report data from the text file
sales_data = []

for line in file_content.split("\n"):
    if line.startswith("ID:"):
        report_details = line.split(", ")
        report_info = {
            "id": int(report_details[0].split(": ")[1]),
            "product_id": int(report_details[1].split(": ")[1]),
            "sales_date": report_details[2].split(": ")[1],
            "quantity_sold": int(report_details[3].split(": ")[1]),
            "total_revenue": float(report_details[4].split(": ")[1]),
        }
        sales_data.append(report_info)

# Extract features (quantity_sold) and target (total_revenue)
X = np.array([report["quantity_sold"] for report in sales_data]).reshape(-1, 1)
y = np.array([report["total_revenue"] for report in sales_data])

# Train a Linear Regression model
model = LinearRegression()
model.fit(X, y)

# Now make predictions on the same dataset (or new data if available)
predictions = model.predict(X)

# Print predictions
print("Predictions:", predictions)

# If you want to visualize the results
plt.scatter(X, y, color='blue', label='Actual data')
plt.plot(X, predictions, color='red', label='Predicted data')
plt.xlabel('Quantity Sold')
plt.ylabel('Total Revenue')
plt.title('Revenue Prediction Based on Quantity Sold')
plt.legend()
plt.show()
