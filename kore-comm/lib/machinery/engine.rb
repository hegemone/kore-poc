module Kore
  module Machinery
    class Engine < Base
      def start
        puts self.test
        #log.debug'Engine#start'
        #EM.run do
        #end
      end
    end
  end
end
