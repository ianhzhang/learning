from pathlib import Path
import hashlib
import os
from PIL import Image
from _ImageDection import ImageContentDector

class DataLoader:
    def loadImageInfo(self, dirpath):
        self.imageDector = ImageContentDector()

        imageList = []
        cnt = 0
        '''
        Search all jpg image under the dirpath
        For each jpg image, find the following:
        * file name
        * file size
        * sha256 key
        * file dimension

        return the list of those information in a list.
        It is a list of map:
        {
            name
            sha256
            size
            width
            height
            exception
        }
        '''
        # Search all files
        # Check if it is a jpeg file (end with jpeg, jpg case insensitive)
        for filename in Path(dirpath).rglob('*.*'):
            filenameStr = filename.as_posix().lower()
            if filenameStr.endswith('.jpeg') or filenameStr.endswith('.jpg'):
                tmpItem = {}
                #tmpItem['name'] = filename.absolute().as_posix()  # absolute path
                tmpItem['name'] = filename.as_posix()    # relative from current dir
                imageList.append(tmpItem)
            else:
                pass

        for index, item in enumerate(imageList):
            #print(index)
            filename = item['name']
            item['exception'] = ''
            try:
                # Get sha-256
                sha256_hash = hashlib.sha256()
                with open(filename,"rb") as f:
                    # Read and update hash string value in blocks of 4K
                    for byte_block in iter(lambda: f.read(4096),b""):
                        sha256_hash.update(byte_block)
                    item['sha256'] = sha256_hash.hexdigest()

                # Get file size
                fileinfo = os.stat(filename)
                item['size'] = fileinfo.st_size

                # Get file dimension
                img = Image.open(filename)
                width, height = img.size
                item['width'] = width
                item['height'] = height

            except Exception as ex:
                print("*****************************")
                item['exception'] = str(ex)
                print(str(ex))
                cnt = cnt + 1

            # ###############################################
            try:
                # Detect Image content
                label, confidence = self.imageDector.detectContent(item['name'])
                item['label'] = label
                item['confidence'] = confidence
            except Exception as ex:
                item['label'] = 'Cannot detect content'
                item['confidence'] =0
        print("total:", len(imageList))
        return imageList
