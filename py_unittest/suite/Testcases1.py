import unittest

class TestCase1(unittest.TestCase):
    def setUp(self):
        print("setUp1")

    def tearDown(self):
        print("tearDown1")

    def test_11(self):
        print("test_11")
        self.assertEqual(True, True)

    def test_12(self):
        print("test_12")
        self.assertEqual(False, True)


if __name__ == '__main__':
    unittest.main()