package main

import (
	"leetcode/array/pkg"
)

func main() {
	// a := [][]int{{1, 2}, {3, 4}}
	// fmt.Println(pkg.MatrixReshape(a, 4, 1))

	// b := []int{9, 77, 63, 22, 92, 9, 14, 54, 8, 38, 18, 19, 38, 68, 58, 19}
	// fmt.Println(pkg.MinSetSize(b))

	// c := [][]int{{1, 2}, {1, 3}}
	// fmt.Println(pkg.KthSmallest(c, 2))

	// d := []int{-9995, -9971, -9957, -9891, -9889, -9865, -9861, -9819, -9792, -9757, -9751, -9733, -9719, -9698, -9680, -9673, -9671, -9663, -9642, -9636, -9596, -9573, -9552, -9533, -9440, -9388, -9347, -9338, -9337, -9295, -9277, -9219, -9183, -9162, -9112, -9071, -9028, -8980, -8975, -8915, -8852, -8824, -8798, -8770, -8762, -8746, -8737, -8701, -8679, -8627, -8618, -8598, -8565, -8543, -8503, -8458, -8378, -8361, -8347, -8236, -8226, -8195, -8126, -8110, -8026, -8013, -8006, -7975, -7940, -7914, -7875, -7873, -7800, -7794, -7770, -7751, -7737, -7713, -7658, -7656, -7542, -7537, -7464, -7412, -7379, -7332, -7327, -7306, -7302, -7110, -7043, -7026, -6933, -6923, -6890, -6855, -6818, -6811, -6776, -6761, -6671, -6603, -6548, -6506, -6501, -6493, -6463, -6420, -6319, -6291, -6289, -6282, -6241, -6180, -6153, -5962, -5931, -5857, -5781, -5719, -5623, -5612, -5607, -5580, -5527, -5505, -5458, -5441, -5386, -5295, -5224, -5185, -5172, -5117, -5108, -5084, -5034, -4974, -4970, -4953, -4907, -4781, -4747, -4712, -4615, -4481, -4465, -4407, -4399, -4364, -4349, -4324, -4322, -4288, -4231, -4229, -4228, -4153, -4040, -4011, -3978, -3965, -3857, -3852, -3748, -3726, -3723, -3671, -3637, -3636, -3632, -3599, -3556, -3525, -3523, -3423, -3409, -3405, -3391, -3357, -3333, -3309, -3246, -3223, -3188, -3140, -3113, -3072, -3062, -2995, -2977, -2970, -2956, -2934, -2925, -2862, -2743, -2696, -2688, -2672, -2652, -2578, -2570, -2561, -2461, -2441, -2423, -2417, -2394, -2296, -2274, -2273, -2267, -2240, -2190, -2180, -2098, -2046, -2039, -2002, -1989, -1884, -1883, -1849, -1785, -1774, -1760, -1647, -1636, -1614, -1609, -1579, -1471, -1447, -1440, -1429, -1425, -1422, -1409, -1355, -1242, -1226, -1217, -1216, -1163, -1147, -1143, -1047, -1014, -915, -910, -876, -851, -834, -818, -810, -770, -683, -607, -557, -534, -527, -511, -502, -498, -439, -270, -251, -226, -132, -116, 38, 94, 132, 178, 218, 226, 246, 282, 358, 371, 404, 426, 445, 450, 477, 485, 506, 530, 579, 637, 669, 728, 751, 784, 818, 838, 863, 888, 953, 955, 960, 967, 983, 1001, 1017, 1087, 1088, 1096, 1162, 1189, 1204, 1222, 1231, 1314, 1326, 1339, 1413, 1471, 1505, 1568, 1585, 1593, 1629, 1677, 1686, 1732, 1752, 1757, 1771, 1786, 1805, 1836, 1870, 1935, 1937, 1990, 2001, 2008, 2043, 2093, 2131, 2158, 2237, 2271, 2278, 2298, 2340, 2350, 2375, 2399, 2418, 2425, 2426, 2427, 2502, 2568, 2602, 2651, 2684, 2721, 2756, 2854, 2969, 2980, 3131, 3135, 3143, 3167, 3222, 3226, 3251, 3256, 3259, 3292, 3306, 3323, 3394, 3495, 3497, 3539, 3549, 3570, 3637, 3719, 3795, 3860, 3894, 3909, 3912, 3914, 3923, 3926, 3947, 3965, 4006, 4010, 4022, 4023, 4044, 4057, 4098, 4143, 4144, 4154, 4187, 4332, 4418, 4564, 4599, 4619, 4660, 4692, 4785, 4931, 5004, 5007, 5042, 5066, 5068, 5086, 5109, 5113, 5152, 5168, 5225, 5234, 5243, 5352, 5416, 5428, 5500, 5533, 5549, 5552, 5556, 5625, 5755, 5798, 5803, 5835, 5836, 5923, 6002, 6008, 6021, 6028, 6040, 6043, 6077, 6117, 6139, 6145, 6163, 6175, 6181, 6245, 6353, 6355, 6362, 6389, 6391, 6400, 6473, 6511, 6566, 6730, 6732, 6782, 6803, 6840, 6942, 6975, 6999, 7061, 7098, 7139, 7204, 7211, 7213, 7244, 7380, 7387, 7392, 7396, 7411, 7510, 7602, 7638, 7660, 7685, 7695, 7778, 7789, 7806, 7816, 7818, 7834, 7906, 7928, 7976, 7997, 8116, 8148, 8149, 8179, 8187, 8201, 8230, 8257, 8276, 8339, 8353, 8384, 8496, 8564, 8575, 8642, 8645, 8653, 8771, 8785, 8809, 8814, 8866, 8873, 8889, 8900, 8942, 9034, 9070, 9090, 9109, 9115, 9190, 9205, 9215, 9217, 9298, 9363, 9449, 9491, 9553, 9595, 9679, 9716, 9762, 9796, 9819, 9820, 9832, 9879, 9925, 9933, 9945, 9888, 9887, 9838, 9828, 9735, 9720, 9594, 9585, 9584, 9580, 9575, 9538, 9429, 9356, 9310, 9297, 9266, 9257, 9227, 9223, 9174, 9143, 9141, 9079, 9035, 9001, 8969, 8962, 8940, 8931, 8776, 8758, 8742, 8736, 8728, 8695, 8625, 8510, 8406, 8377, 8277, 8220, 8138, 8026, 7929, 7911, 7904, 7731, 7717, 7670, 7650, 7577, 7553, 7526, 7475, 7455, 7409, 7326, 7313, 7292, 7285, 7253, 7182, 7027, 7008, 6993, 6895, 6866, 6863, 6800, 6770, 6714, 6681, 6526, 6480, 6459, 6437, 6377, 6321, 6303, 6072, 5981, 5954, 5903, 5843, 5801, 5785, 5783, 5782, 5745, 5709, 5662, 5579, 5575, 5561, 5516, 5512, 5397, 5393, 5373, 5369, 5357, 5339, 5238, 5202, 5119, 5096, 4990, 4951, 4933, 4930, 4914, 4909, 4866, 4840, 4826, 4805, 4771, 4724, 4611, 4539, 4538, 4480, 4415, 4369, 4294, 4290, 4253, 4164, 3983, 3930, 3844, 3811, 3768, 3675, 3611, 3516, 3513, 3500, 3464, 3463, 3446, 3434, 3393, 3378, 3352, 3348, 3286, 3214, 3192, 3178, 3172, 3101, 3058, 2996, 2873, 2871, 2857, 2777, 2766, 2750, 2708, 2634, 2614, 2484, 2378, 2336, 2323, 2198, 2160, 2108, 2047, 2014, 1952, 1949, 1882, 1858, 1843, 1833, 1707, 1684, 1647, 1637, 1613, 1583, 1572, 1543, 1315, 1202, 1135, 1119, 1089, 1076, 988, 901, 864, 848, 847, 832, 831, 698, 695, 645, 602, 592, 568, 515, 510, 382, 347, 247, 228, 216, 151, -32, -152, -174, -268, -352, -375, -393, -442, -445, -464, -545, -760, -906, -962, -964, -1109, -1114, -1165, -1178, -1243, -1271, -1404, -1476, -1482, -1495, -1515, -1592, -1617, -1626, -1657, -1688, -1740, -1790, -1826, -1851, -1920, -1938, -1939, -1966, -1991, -2056, -2183, -2188, -2245, -2246, -2257, -2259, -2309, -2319, -2383, -2387, -2426, -2482, -2541, -2562, -2592, -2619, -2655, -2720, -2764, -2828, -2947, -2949, -2962, -3046, -3067, -3087, -3095, -3146, -3181, -3186, -3202, -3327, -3479, -3509, -3596, -3609, -3611, -3708, -3744, -3754, -3761, -3775, -3862, -3879, -3987, -4116, -4133, -4141, -4377, -4486, -4541, -4582, -4607, -4625, -4682, -4687, -4733, -4745, -4758, -4842, -4864, -5031, -5057, -5109, -5112, -5123, -5125, -5132, -5269, -5304, -5430, -5455, -5460, -5463, -5487, -5489, -5515, -5531, -5579, -5586, -5668, -5683, -5692, -5754, -5784, -5859, -5866, -5910, -5924, -5999, -6089, -6160, -6186, -6198, -6246, -6274, -6354, -6383, -6408, -6448, -6505, -6542, -6565, -6574, -6630, -6773, -6792, -6801, -6821, -6857, -6877, -7008, -7074, -7108, -7149, -7188, -7226, -7282, -7284, -7337, -7343, -7352, -7368, -7466, -7496, -7505, -7594, -7626, -7711, -7738, -7783, -7789, -7812, -7818, -7835, -7844, -7847, -7849, -7856, -7889, -7894, -7917, -8037, -8063, -8064, -8103, -8107, -8161, -8262, -8303, -8324, -8336, -8359, -8440, -8469, -8475, -8478, -8481, -8491, -8515, -8577, -8635, -8713, -8773, -8797, -8884, -8940, -8943, -8982, -9002, -9030, -9032, -9113, -9158, -9189, -9276, -9371, -9560, -9598, -9603, -9616, -9622, -9646, -9654, -9695, -9750, -9762, -9795, -9810, -9876}
	// fmt.Println(pkg.FindPeakElement(d))

	// e := []int{1, 2, 3, 4, 5}
	// s := pkg.Constructor(e)
	// fmt.Println(s.Shuffle())

	n1 := []*pkg.ListNode{
		{ Val: 1, Next: &pkg.ListNode{ Val: 4, Next: &pkg.ListNode{ Val: 5, Next: nil }}},
		{ Val: 1, Next: &pkg.ListNode{ Val: 3, Next: &pkg.ListNode{ Val: 4, Next: nil }}},
		nil, nil,
		{ Val: 2, Next: &pkg.ListNode{ Val: 6, Next: nil }},	
	}
	n := pkg.MergeKLists(n1)
	pkg.PrintNode(n)
}
