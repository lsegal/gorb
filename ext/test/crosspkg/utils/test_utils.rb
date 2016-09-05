require_relative './utils'
require_relative '../data/data'

include Test::Crosspkg
include Test::Crosspkg::Data

rgb = RGB.new
rgb.r = 20
rgb.g = 40
rgb.b = 60
p rgb

hsv = Test::Crosspkg::Utils.to_hsv(rgb)
puts hsv
p hsv
