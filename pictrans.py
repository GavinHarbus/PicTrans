from PIL import Image
import argparse as apa

def cleanPic(pic_name):
    try:
        img = Image.open("./static/pics/"+pic_name+".jpg")
        out = img.point(lambda x:255 if x>75 else 0) 
        out.convert('L').save("./static/pics/res"+pic_name+".jpg")
        print("true")
    except Exception:
        print("false")

if __name__ == '__main__':
    parser = apa.ArgumentParser(prog="convert")
    parser.add_argument("--input", required=True, help="The xml file path",type=str)
    args = parser.parse_args()
    cleanPic(args.input)

