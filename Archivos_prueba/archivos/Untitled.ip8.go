{
 "cells": [
  {
   "cell_type": "code",
   "execution_count": 2,
   "metadata": {},
   "outputs": [],
   "source": [
    "import shutil\n",
    "from os import listdir\n",
    "from os.path import isfile, join"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 5,
   "metadata": {},
   "outputs": [
    {
     "ename": "FileNotFoundError",
     "evalue": "[WinError 3] El sistema no puede encontrar la ruta especificada: ''",
     "output_type": "error",
     "traceback": [
      "\u001b[1;31m---------------------------------------------------------------------------\u001b[0m",
      "\u001b[1;31mFileNotFoundError\u001b[0m                         Traceback (most recent call last)",
      "\u001b[1;32m<ipython-input-5-5d18494bd7f5>\u001b[0m in \u001b[0;36m<module>\u001b[1;34m\u001b[0m\n\u001b[0;32m      1\u001b[0m \u001b[0mmypath\u001b[0m \u001b[1;33m=\u001b[0m \u001b[1;34m\"\"\u001b[0m\u001b[1;33m\u001b[0m\u001b[1;33m\u001b[0m\u001b[0m\n\u001b[1;32m----> 2\u001b[1;33m \u001b[0monlyfiles\u001b[0m \u001b[1;33m=\u001b[0m \u001b[1;33m[\u001b[0m\u001b[0mf\u001b[0m \u001b[1;32mfor\u001b[0m \u001b[0mf\u001b[0m \u001b[1;32min\u001b[0m \u001b[0mlistdir\u001b[0m\u001b[1;33m(\u001b[0m\u001b[0mmypath\u001b[0m\u001b[1;33m)\u001b[0m \u001b[1;32mif\u001b[0m \u001b[0misfile\u001b[0m\u001b[1;33m(\u001b[0m\u001b[0mjoin\u001b[0m\u001b[1;33m(\u001b[0m\u001b[0mmypath\u001b[0m\u001b[1;33m,\u001b[0m \u001b[0mf\u001b[0m\u001b[1;33m)\u001b[0m\u001b[1;33m)\u001b[0m\u001b[1;33m]\u001b[0m\u001b[1;33m\u001b[0m\u001b[1;33m\u001b[0m\u001b[0m\n\u001b[0m",
      "\u001b[1;31mFileNotFoundError\u001b[0m: [WinError 3] El sistema no puede encontrar la ruta especificada: ''"
     ]
    }
   ],
   "source": [
    "mypath = \"\"\n",
    "onlyfiles = [f for f in listdir(mypath) if isfile(join(mypath, f))]"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": null,
   "metadata": {},
   "outputs": [],
   "source": []
  }
 ],
 "metadata": {
  "kernelspec": {
   "display_name": "Python 3",
   "language": "python",
   "name": "python3"
  },
  "language_info": {
   "codemirror_mode": {
    "name": "ipython",
    "version": 3
   },
   "file_extension": ".py",
   "mimetype": "text/x-python",
   "name": "python",
   "nbconvert_exporter": "python",
   "pygments_lexer": "ipython3",
   "version": "3.8.3"
  }
 },
 "nbformat": 4,
 "nbformat_minor": 4
}
