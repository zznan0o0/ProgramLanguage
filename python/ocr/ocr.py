import io
from PIL import Image as PI
import pyocr
import pyocr.builders
from cnocr import CnOcr

print( pyocr.get_available_tools())


tool = pyocr.get_available_tools()[0]
img_url = "ocr.jpg"
with open(img_url, 'rb') as f:
    a = f.read()
new_img = PI.open(io.BytesIO(a))

# 金额
left = 741
top = 420
right = 850
bottom = 445
image_text1 = new_img.crop((left, top, right, bottom))
#展示图片
# image_text1.show()


# 名称
left = 155
top = 450
right = 450
bottom = 470
image_obj2 = new_img.crop((left, top, right, bottom))
# image_obj2.show()

#纳税人识别号
left = 155
top = 470
right = 450
bottom = 490
image_text3 = new_img.crop((left, top, right, bottom))
#展示图片
image_text3.show()

# 开票人
left = 528
top = 550
right = 670
bottom = 600
image_obj4 = new_img.crop((left, top, right, bottom))
image_obj4.show()



image_obj2.save("tmp.jpg")
ocr = CnOcr()
res = ocr.ocr("tmp.jpg")
print("".join(res[0]))


image_obj4.save("tmp.jpg")
ocr = CnOcr()
res = ocr.ocr("tmp.jpg")
print("".join(res[0]))