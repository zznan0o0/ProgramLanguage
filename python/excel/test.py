from openpyxl import load_workbook
from openpyxl_image_loader import SheetImageLoader

# 定义需要处理的excel路径和sheetname
excel_root = "./1.xlsx"
sheet_name = "0"

workbook = load_workbook(filename=excel_root)
sheet = workbook[sheet_name]

image_loader = SheetImageLoader(sheet)
print(image_loader._images)

# name对应单元格, 例如b2等; value就是对于的pilimage对象
for name, value in image_loader._images.keys():
    print(name)
    print()
    value.save