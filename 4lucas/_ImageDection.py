import torch
from torchvision import models
from torchvision import transforms
from PIL import Image

# Use pytorch pre-trained model (resnet101) to identify the contents in a image
class ImageContentDector:
    def __init__(self):
        self.enabled = True
        try:
            self.resnet = models.resnet101(pretrained=True)

            self.preprocess = transforms.Compose([
                transforms.Resize(256),
                transforms.CenterCrop(224),
                transforms.ToTensor(),
                transforms.Normalize(
                mean=[0.485, 0.456, 0.406],
                std=[0.229, 0.224, 0.225]
            )])
            with open('./imagenet_classes.txt') as f:
                self.labels = [line.strip() for line in f.readlines()]

            self.resnet.eval()
        except Exception as ex:
            self.enabled = False

    def detectContent(self, filename):
        if self.enabled == False:
            return "", 0.0

        img = Image.open(filename)
        img_t = self.preprocess(img)
        batch_t = torch.unsqueeze(img_t, 0)

        out = self.resnet(batch_t)
        _, index = torch.max(out, 1)

        percentage = torch.nn.functional.softmax(out, dim=1)[0] * 100

        return self.labels[index[0]], round(percentage[index[0]].item(),2)