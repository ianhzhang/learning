import unittest
from Testcases1 import TestCase1
from Testcases2 import TestCase2

if __name__ == '__main__':
    suite = unittest.TestSuite()

    suite.addTest(unittest.makeSuite(TestCase1))
    suite.addTest(unittest.makeSuite(TestCase2))

    runner = unittest.TextTestRunner()
    print(runner.run(suite))