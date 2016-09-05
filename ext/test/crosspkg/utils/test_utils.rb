require_relative './utils'
require_relative '../data/data'

include Test::Crosspkg
include Test::Crosspkg::Data

class Color
  def initialize(hsv, rgb)
    self.hsv = hsv
    self.rgb = rgb
  end

  def inspect
    "#{hsv} / RGB(r=#{rgb.r} g=#{rgb.g} b=#{rgb.b})"
  end
end

rgb = RGB.new
rgb.r = 20
rgb.g = 40
rgb.b = 60
p rgb

hsv = Test::Crosspkg::Utils.to_hsv(rgb)
puts hsv
p hsv

p Color.new(hsv, rgb)
