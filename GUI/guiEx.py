import tkinter as tk
from tkinter import *
from PIL import Image, ImageTk

root = tk.Tk()
window = root.title('БАСТИОН')
root.iconbitmap("D:\code\Hakatonich\img\img.ico")


# основа
mainarea = tk.Frame(root, bg='#CCC', width=500, height=500)
mainarea.pack(expand=True, fill='both', side='right')

def Termo_Window():

    pass


# Херня с боку
sidebar = tk.Frame(root, width=40, bg='white', height=500, relief='sunken', borderwidth=2)
sidebar.pack(expand=False, fill='both', side='left', anchor='nw')

image1 = ImageTk.PhotoImage(Image.open('D:\code\Hakatonich\img\Termo.png').resize((40,40)))
btn1 = Button(sidebar, image=image1, bg='white',relief='flat', command=Termo_Window)
btn1.pack(side=TOP)

image2 = ImageTk.PhotoImage(Image.open('D:\code\Hakatonich\img\Light.png').resize((40,45)))
btn2 = Button(sidebar, image=image2, bg='white',relief='flat')
btn2.pack(side=TOP)

image3 = ImageTk.PhotoImage(Image.open('D:\code\Hakatonich\img\Whater.png').resize((40,45)))
btn3 = Button(sidebar, image=image3, bg='white',relief='flat')
btn3.pack(side=TOP)







root.mainloop()