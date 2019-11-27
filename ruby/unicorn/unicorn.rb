class Unicorn
  attr_reader :name
  
  def initialize(name, color="white")
    @name = name
    @color = color
  end

  def name
    @name
  end

  def color
    @color
  end

  def white?
    color == "white"
  end

  def say(something)
     "**;* #{something} **;*"
  end
end
