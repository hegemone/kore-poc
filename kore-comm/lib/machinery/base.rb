module Kore
  module Machinery
    #include Inject['log', 'config']
    class Base
      attr_reader :test
      def initialize
        @test = 'test'
        #@log = d[:log]
        #@config = d[:config]
    end
  end
end
