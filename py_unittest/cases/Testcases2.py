import unittest

class TestCase2(unittest.TestCase):
    def setUp(self):
        print("setUp")

    def tearDown(self):
        print("tearDown")

    def test_11(self):
        print("test_21")
        self.assertEqual(True, True)

    def test_2(self):
        print("test_22")
        self.assertEqual(True, True)

if __name__ == '__main__':
    unittest.main()